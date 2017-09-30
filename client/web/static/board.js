goog.provide('ticktacktoe.Board');
goog.require('proto.Player');

/**
 * @constructor
 * @param{Phaser.Game} game
 * @param{number} cellSize
 */
function Board(game, cellSize) {
  this.game = game;
  this.cellSize = cellSize;
  this.cells = [
    [false, false, false],
    [false, false, false],
    [false, false, false],
  ];
}

ticktacktoe.Board = Board;

/**
 * @param {proto.Player} player
 * @return {string}
 */
Board.color = function(player) {
  if (player == proto.Player.A) {
    return 'blue';
  }
  if (player == proto.Player.B) {
    return 'red';
  }
  console.log('unknown player value', player);
  return 'black';
};

/**
 * @return {void}
 */
Board.prototype.onPhaserCreate = function() {
  // Add a grid board to the UI.
  this.game.create.grid('board',
                        this.cellSize * 3 + 1,
                        this.cellSize * 3 + 1,
                        this.cellSize,
                        this.cellSize,
                        '#000000');
  this.game.add.sprite(0, 0, 'board');

  // Setup bitmap, used for marking the board.
  this.canvas = this.game.make.bitmapData(this.cellSize * 3,
                                          this.cellSize * 3);
  this.canvas.addToWorld(0, 0);
};

/**
 * @param {number} row
 * @param {number} col
 * @return {boolean}
 */
Board.prototype.canTake = function(row, col) {
  return !this.cells[row][col];
};

/**
 * @param {number} row
 * @param {number} col
 * @param {proto.Player} player
 * @return {void}
 */
Board.prototype.update = function(row, col, player) {
  this.cells[row][col] = true;
  this.mark(row, col, Board.color(player));
};


/**
 * @param {number} row
 * @param {number} col
 * @param {string} color
 * @return {void}
 */
Board.prototype.mark = function(row, col, color) {
  this.canvas.circle((col + 0.5) * this.cellSize,
                     (row + 0.5) * this.cellSize,
                     0.4 * this.cellSize,
                     color);
  this.canvas.circle((col + 0.5) * this.cellSize,
                     (row + 0.5) * this.cellSize,
                     0.35 * this.cellSize,
                     '#ffffff');
};
