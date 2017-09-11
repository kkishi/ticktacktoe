goog.require('goog.crypt.base64');
goog.require('proto.Join');
goog.require('proto.Move');
goog.require('proto.Request');
goog.require('proto.Response');

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
    { create: this.phaserCreate.bind(this) });
  this.taken = [
    [false, false, false],
    [false, false, false],
    [false, false, false],
  ];

  // Set up WebSocket.
  this.connection = new WebSocket(
    'ws://' + location.host.replace(":8081", ":8080") + "/ttt");
  this.connection.onopen = this.socketOnopen.bind(this);
  this.connection.error = this.socketError.bind(this);
  this.connection.onmessage = this.socketOnmessage.bind(this);

  this.state = TickTackToe.State.SETUP;
}

TickTackToe.Colors = [
  'blue',
  'red'
];

TickTackToe.Players = {
  SELF: 0,
  OPPONENT: 1
};

TickTackToe.State = {
  UNKNOWN: 0,
  SETUP: 1,
  MAKE_MOVE: 2,
  WAIT_MOVE: 3,
  FINISHED: 4
};

/**
 * void
 */
TickTackToe.prototype.phaserCreate = function() {
  this.game.stage.backgroundColor = '#ffffff';

  // Add a grid board to the UI.
  this.game.create.grid('board',
                        this.cellSize * 3 + 1,
                        this.cellSize * 3 + 1,
                        this.cellSize,
                        this.cellSize,
                        '#000000');
  this.game.add.sprite(0, 0, 'board');

  // Register mouse events.
  this.game.input.mouse.capture = true;
  this.game.input.onDown.add(this.onDown.bind(this));

  // Setup bitmap, used for marking the board.
  this.canvas = this.game.make.bitmapData(this.cellSize * 3,
                                          this.cellSize * 3);
  this.canvas.addToWorld(0, 0);
};

/**
 * @param {number} row
 * @param {number} col
 * @param {string} color
 */
TickTackToe.prototype.mark = function(row, col, color) {
  this.canvas.circle((col + 0.5) * this.cellSize,
                     (row + 0.5) * this.cellSize,
                     0.4 * this.cellSize,
                     color);
  this.canvas.circle((col + 0.5) * this.cellSize,
                     (row + 0.5) * this.cellSize,
                     0.35 * this.cellSize,
                     '#ffffff');
};

/**
 * @param {Point} pointer
 */
TickTackToe.prototype.onDown = function(pointer) {
  var row = Math.floor(pointer.y / this.cellSize);
  var col = Math.floor(pointer.x / this.cellSize);
  console.log('clicked', row, col);
  if (this.state != TickTackToe.State.MAKE_MOVE) {
    console.log('not read to take; state', this.state);
    return;
  }
  if (!this.canTake(row, col)) {
    console.log('can not take the cell');
    return;
  }
  this.take(row, col, TickTackToe.Players.SELF);
};

/**
 * @param {number} row
 * @param {number} col
 * @return {boolean}
 */
TickTackToe.prototype.canTake = function(row, col) {
  return !this.taken[row][col];
};

/**
 * @param {number} row
 * @param {number} col
 * @param {number} player
 */
TickTackToe.prototype.take = function(row, col, player) {
  this.taken[row][col] = true;
  this.mark(row, col, TickTackToe.Colors[player]);
  if (player == TickTackToe.Players.OPPONENT) {
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
 * void
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
  var u = goog.crypt.base64.decodeStringToUint8Array(e.data);
  var r = proto.Response.deserializeBinary(u.buffer);
  console.log("response", e.data, u, r.toString());
  var f = r.getFinish();
  if (f) {
    var o = f.getOpponent();
    if (o) {
      this.take(o.getRow(), o.getCol(), TickTackToe.Players.OPPONENT);
    }
    this.state = TickTackToe.State.FINISHED;
    this.connection.close();
    console.log("finished", f.getResult().toString());
    return;
  }
  var mm = r.getMakeMove();
  if (mm) {
    var o = mm.getOpponent();
    if (o) {
      this.take(o.getRow(), o.getCol(), TickTackToe.Players.OPPONENT);
      console.log("opponent took", o.getRow(), o.getCol());
    }
  }
  this.state = TickTackToe.State.MAKE_MOVE;
};

(function() {
  var ttt = new TickTackToe();
})();
