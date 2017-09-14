/**
 * @externs
 * @record
 */
var Phaser = {}

/**
 * @constructor
 */
Phaser.Game = function(a, b, c, d, e) {}
Phaser.Game.prototype.game.add.sprite;
Phaser.Game.prototype.game.add;
Phaser.Game.prototype.game.create;
Phaser.Game.prototype.game.grid;
Phaser.Game.prototype.game.input.mouse.capture;
Phaser.Game.prototype.game.input.mouse;
Phaser.Game.prototype.game.input.onDown.add;
Phaser.Game.prototype.game.input.onDown;
Phaser.Game.prototype.game.input;
Phaser.Game.prototype.game.make.bitmapData = function() {};
Phaser.Game.prototype.game.make.bitmapData.prototype.addToWorld;
Phaser.Game.prototype.game.make.bitmapData.prototype.circle;
Phaser.Game.prototype.game.make;
Phaser.Game.prototype.game.stage.backgroundColor;
Phaser.Game.prototype.game.stage;
Phaser.Game.prototype.game;

Phaser.AUTO;
Phaser.CANVAS;

/**
 * @constructor
 *
 * HACK: Ideally this should also have @externs annotation, but that causes a
 * compiler warning because Closure compiler somehow removes this type
 * information before type checking, seemingly because Point is only used in a
 * function signature type assertion.
 */
function Point() {}
Point.prototype.x;
Point.prototype.y;
