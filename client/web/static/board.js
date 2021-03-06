goog.provide('ticktacktoe.Board');
goog.require('proto.Player');

/**
 * @param {!Phaser.Game} game
 * @param {number} cellSize
 * @constructor
 */
ticktacktoe.Board = function(game, cellSize) {
  this.game = game;
  this.cellSize = cellSize;
  this.cells = [
    [false, false, false],
    [false, false, false],
    [false, false, false],
  ];

  /** @type {?Phaser.BitmapData} */
  this.canvas = null;
};

/**
 * @param {proto.Player} player
 * @return {string}
 */
ticktacktoe.Board.color = function(player) {
  if (player == proto.Player.A) {
    return 'blue';
  }
  if (player == proto.Player.B) {
    return 'red';
  }
  console.log('unknown player value', player);
  return 'black';
};

/** @return {void} */
ticktacktoe.Board.prototype.onPhaserCreate = function() {
  // Add a grid board to the UI.
  this.game.create.grid('board',
                        this.cellSize * 3 + 1,
                        this.cellSize * 3 + 1,
                        this.cellSize,
                        this.cellSize,
                        '#000000',
                        true,
                        (function() {
                          this.game.add.sprite(0, 0, 'board');
                        }).bind(this));

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
ticktacktoe.Board.prototype.canTake = function(row, col) {
  return !this.cells[row][col];
};

/**
 * @param {number} row
 * @param {number} col
 * @param {proto.Player} player
 */
ticktacktoe.Board.prototype.update = function(row, col, player) {
  this.cells[row][col] = true;
  this.mark(row, col, ticktacktoe.Board.color(player));
};


/**
 * @param {number} row
 * @param {number} col
 * @param {string} color
 */
ticktacktoe.Board.prototype.mark = function(row, col, color) {
  this.canvas.circle((col + 0.5) * this.cellSize,
                     (row + 0.5) * this.cellSize,
                     0.4 * this.cellSize,
                     color);
  this.canvas.circle((col + 0.5) * this.cellSize,
                     (row + 0.5) * this.cellSize,
                     0.35 * this.cellSize,
                     '#ffffff');
};
