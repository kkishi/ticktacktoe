goog.require('goog.crypt.base64');
goog.require('proto.Join');
goog.require('proto.Move');
goog.require('proto.Request');
goog.require('proto.Response');
goog.require('ticktacktoe.Board');

/**
 * @constructor
 */
function TickTackToe() {
  // Set up Phaser.
  this.cellSize = 200;
  this.game = new Phaser.Game(
    this.cellSize * 3 + 1,
    this.cellSize * 3 + 1,
    // TODO: Switch to AUTO. Painting bitmap somehow doesn't work with WebGL.
    Phaser.CANVAS,
    'phaser-game',
    { create: this.onPhaserCreate.bind(this) });
  this.board = new ticktacktoe.Board(this.game, this.cellSize);

  // Set up WebSocket.
  this.connection = new WebSocket(
    'ws://' + location.host.replace(":8081", ":8080") + "/ttt");
  this.connection.onopen = this.socketOnopen.bind(this);
  this.connection.error = this.socketError.bind(this);
  this.connection.onmessage = this.socketOnmessage.bind(this);

  this.state = TickTackToe.State.SETUP;
}

TickTackToe.State = {
  UNKNOWN: 0,
  SETUP: 1,
  MAKE_MOVE: 2,
  WAIT_MOVE: 3,
  FINISHED: 4
};

/**
 * @return {void}
 */
TickTackToe.prototype.onPhaserCreate = function() {
  this.game.stage.backgroundColor = '#ffffff';

  // Register mouse events.
  this.game.input.mouse.capture = true;
  this.game.input.onDown.add(this.onDown.bind(this));

  this.board.onPhaserCreate();
};

/**
 * @param {Point} pointer
 * @return {void}
 */
TickTackToe.prototype.onDown = function(pointer) {
  var row = Math.floor(pointer.y / this.cellSize);
  var col = Math.floor(pointer.x / this.cellSize);
  console.log('clicked', row, col);
  if (this.state != TickTackToe.State.MAKE_MOVE) {
    console.log('not read to take; state', this.state);
    return;
  }
  if (!this.board.canTake(row, col)) {
    console.log('can not take the cell');
    return;
  }
  var m = new proto.Move;
  m.setRow(row);
  m.setCol(col);
  var r = new proto.Request;
  r.setMove(m);
  var b = goog.crypt.base64.encodeByteArray(r.serializeBinary());
  this.connection.send(b);
  this.state = TickTackToe.State.WAIT_MOVE;
};

/**
 * @return {void}
 */
TickTackToe.prototype.socketOnopen = function() {
  var j = new proto.Join;
  j.setName("Tick");
  var r = new proto.Request;
  r.setJoin(j);
  var s = r.serializeBinary();
  var b = goog.crypt.base64.encodeByteArray(s);
  console.log("request", r.toString(), s, b);
  this.connection.send(b);
};

/**
 * @param {string} error
 */
TickTackToe.prototype.socketError = function(error) {
  console.log('websocket error: ' + error);
  this.connection.close();
  this.state = TickTackToe.State.FINISHED;
};

/**
 * @param {Object} e
 */
TickTackToe.prototype.socketOnmessage = function(e) {
  var b = goog.crypt.base64.decodeStringToUint8Array(e.data);
  var r = proto.Response.deserializeBinary(b.buffer);
  console.log("response", e.data, b, r.toString(), r.toObject());
  var f = r.getFinish();
  if (f) {
    this.state = TickTackToe.State.FINISHED;
    this.connection.close();
    console.log("finished", f.getResult().toString());
    return;
  }
  var u = r.getUpdate();
  if (u) {
    this.board.update(u.getRow(), u.getCol(), u.getPlayer());
    console.log("cell taken", u.getRow(), u.getCol(), u.getPlayer());
    return;
  }
  var mm = r.getMakeMove();
  if (mm) {
    this.state = TickTackToe.State.MAKE_MOVE;
    return;
  }
};

(function() {
  var ttt = new TickTackToe();
})();
