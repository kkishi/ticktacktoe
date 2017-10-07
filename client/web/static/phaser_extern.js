/**
 * @fileoverview Externs file for Phaser.
 * @externs
 */

/**
 * @type {Object}
 * @const
 */
var Phaser;

/** @constructor */
Phaser.Game = function(a, b, c, d, e) {
  /** @type {Phaser.GameObjectFactory} */
  this.add;

  /** @type {Phaser.Create} */
  this.create;

  /** @type {Phaser.Input} */
  this.input;

  /** @type {Phaser.GameObjectCreator} */
  this.make;

  /** @type {Phaser.Stage} */
  this.stage;
}

/** @constructor */
Phaser.GameObjectFactory = function() {};

/** @return {void} */
Phaser.GameObjectFactory.prototype.sprite = function(a, b, c) {};

/** @constructor */
Phaser.Create = function() {}

/** @return {void} */
Phaser.Create.prototype.grid = function(a, b, c, d, e, f, g, h) {};

/** @constructor */
Phaser.Input = function() {
  /** @type {Phaser.Mouse} */
  this.mouse;

  /** @type {Phaser.Signal} */
  this.onDown;
}

/** @constructor */
Phaser.GameObjectCreator = function() {};

/** @return {Phaser.BitmapData} */
Phaser.GameObjectCreator.prototype.bitmapData = function(a, b) {};

/** @constructor */
Phaser.BitmapData = function() {};

/** @return {void} */
Phaser.BitmapData.prototype.addToWorld = function(a, b) {};

/** @return {void} */
Phaser.BitmapData.prototype.circle = function(a, b, c, d) {};

/** @constructor */
Phaser.Stage = function() {
  /** @type {string} */
  this.backgroundColor;
};

/** @constructor */
Phaser.Mouse = function() {};

/** @type {boolean} */
Phaser.Mouse.prototype.capture;

/** @constructor */
Phaser.Signal = function() {};

/** @return {void} */
Phaser.Signal.prototype.add = function(a) {};

/** @type {number} */
Phaser.AUTO;

/** @type {number} */
Phaser.CANVAS;

/** @constructor */
function Point() {
  /** @type {number} */
  this.x;

  /** @type {number} */
  this.y;
}
