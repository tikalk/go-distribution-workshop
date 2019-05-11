// TODO remove dead players from display

let Player = class {
	constructor(model, config) {
		this.config = config;
		this.clear = true;			// do I have enough space around me?

		this.container = new createjs.Container();
		this.targetLabelAlpha = 1;

		this.hl = new createjs.Shape();
		this.hl.graphics.ss(8).s("white").dc(0, 0, config.playerRadius).es();
		this.hl.alpha = 0.8;
		this.container.addChild(this.hl);

		this.graphic = new createjs.Shape();
		this.graphic.graphics.f(model.team_id).dc(0, 0, config.playerRadius).ef();

		let that = this;
		this.graphic.addEventListener("mouseover", function(){
			that.mouseover = true;
		});

		this.graphic.addEventListener("mouseout", function(){
			that.mouseover = false;
		});

		this.container.addChild(this.graphic);

		this.hl.visible = false;

		this.label = new createjs.Text(model.item_label, "15px Arial", "white");
		this.label.textAlign = "center";
		this.label.width = 200;
		this.label.x = 0;
		this.label.y = config.playerRadius * 3;
		this.label.textBaseline = "alphabetic";
		this.container.addChild(this.label);

		this.update(model, true);

		config.stage.addChild(this.container);
	}

	update(model, hard){
		if(model) {
			this.targetX = this.config.maxWidth * model.x / 100;
			this.targetY = this.config.maxHeight * model.y / 100;
			this.clear = true;
		} else {
			if (hard) {
				this.container.x = this.targetX;
				this.container.y = this.targetY;
			}
			else {
				this.container.x += (this.targetX - this.container.x) * 0.2;
				this.container.y += (this.targetY - this.container.y) * 0.2;
			}
			this.targetLabelAlpha = this.clear ? 1 : 0;
			if (this.hl.visible || this.mouseover){
				this.label.alpha = 1;
			} else {
				this.label.alpha += (this.targetLabelAlpha - this.label.alpha) * 0.2
			}

		}
	}

	highlight(){
		this.hl.visible = true;
	}

	forget(){
		this.hl.visible = false;
	}

	kill(){
		// TODO add some cool animation
		this.config.stage.removeChild(this.container)
	}

	getTarget(p){
		return {x: this.targetX, y: this.targetY};
	}

	targetDistanceTo(other){
		let otherTarget = other.getTarget();
		let dx = this.targetX - otherTarget.x;
		let dy = this.targetY - otherTarget.y;
		return Math.sqrt(Math.pow(dx, 2) + Math.pow(dy, 2));
	}
}

let Ball = class {
	constructor(model, config) {
		this.config = config;

		this.graphic = new createjs.Shape();
		this.graphic.graphics.f("white").dc(0, 0, config.ballRadius).ef();

		this.shadow = new createjs.Shape();
		this.shadow.graphics.f("black").de(-config.ballRadius * 0.3, config.ballRadius * 0.6, config.ballRadius * 2, config.ballRadius).ef();
		this.shadow.alpha = 0.2;

		this.container = new createjs.Container();
		this.container.addChild(this.shadow);
		this.container.addChild(this.graphic);

		if(!model){
			model = {x: 0, y : 0}
		}
		this.update(model, true);
		config.stage.addChild(this.container);
	}

	update(model, hard){
		if(model) {
			this.targetX = this.config.maxWidth * model.x / 100;
			this.targetY = this.config.maxHeight * model.y / 100;
		} else {

			if (hard) {
				this.container.x = this.targetX;
				this.container.y = this.targetY;
			}
			else {
				this.container.x += (this.targetX - this.container.x) * 0.2;
				this.container.y += (this.targetY - this.container.y) * 0.2;
			}
		}
	}

}

var config = {
	maxWidth: window.innerWidth,
	maxHeight: window.innerHeight,
	playerRadius: 10,
	displayInterval: 100,
	animationInterval: 25,
	ballRadius: 5,
	playerLabelRadius: 100,
}

async function init() {
	let canvas = document.getElementById("main_canvas");
	canvas.width = window.innerWidth;
	canvas.height = window.innerHeight;



	let stage = new createjs.Stage("main_canvas");
	drawField(stage);
	stage.enableMouseOver(10);

	config.stage = stage;

	let players = [];
	let ball = new Ball(null, config);
	let lastHolder = null;

	async function getDisplay(){
		let res = await fetch("/display");
		let display = await res.json();

		let updatedKeys = {}


		for (let key in display.items){
			let item = display.items[key];
			key = "player|" + item.team_id + "|" + item.item_id;
			switch(item.item_type){
				case "player":
					updatedKeys[key] = true;
					if (key in players) {
						players[key].update(item);
					} else {
						players[key] = new Player(item, config);
					}
					break;
				case "ball":
					ball.update(item);
					if(item.item_id && players[key]) {


						if(lastHolder){
							lastHolder.forget();
						}
						lastHolder = players[key]
						lastHolder.highlight();
					}
					break;
			}
		}
		lastHolder = cleanupDeadPlayers(players, updatedKeys, lastHolder);
		prettifyDisplay(players)
	}

	function animate() {
		for (let i in players){
			players[i].update();
		}

		ball.update();

		stage.update();
	}

	setInterval(getDisplay, config.displayInterval)
	setInterval(animate, config.animationInterval)

}

function drawField(stage){
	let field = new createjs.Shape();
	let g = field.graphics;
	let baseColor1 = "#060";
	let baseColor2 = "#090";
	let hlColor = "rgba(0, 180, 0, 0.5)";
	let lineColor = "#FFF";
	let borderGap = 0.05;
	let centerCircleRadius = 0.12;
	let gateFarWidth = 0.1;
	let gateFarHeight = 0.25;
	let gateNearWidth = 0.04;
	let gateNearHeight = 0.12;
	let gateWidth = 0.015;
	let gateHeight = 0.05;
	let lengthDivisions = 14;

	// Base
	g.lf([baseColor1,baseColor2], [0, 1], config.maxWidth, 0, 0, config.maxHeight).dr(0, 0, config.maxWidth, config.maxHeight).ef();

	// Divisions
	for (let i = 0; i < lengthDivisions; i++){
		if(i %2 == 0) {
			g.f(hlColor).dr(i * config.maxWidth / lengthDivisions, 0, config.maxWidth / lengthDivisions, config.maxHeight).ef();
		}
	}

	// Border
	g.ss(2).s(lineColor).dr(config.maxHeight * borderGap, config.maxHeight * borderGap, config.maxWidth - (config.maxHeight * 2 * borderGap), config.maxHeight * (1 - 2 * borderGap)).es();

	// Center
	g.ss(2).s(lineColor).mt(config.maxWidth / 2, 0).lt(config.maxWidth / 2, config.maxHeight).es();
	g.ss(2).s(lineColor).dc(config.maxWidth / 2, config.maxHeight / 2, config.maxHeight * centerCircleRadius).es();

	// Gates
	g.ss(2).s(lineColor).dr(config.maxHeight * borderGap, config.maxHeight * (0.5 - gateFarHeight),	config.maxWidth * gateFarWidth,	config.maxHeight * 2 * gateFarHeight).es();
	g.ss(2).s(lineColor).dr(config.maxWidth * (1 - gateFarWidth) - config.maxHeight * (borderGap), config.maxHeight * (0.5 - gateFarHeight), config.maxWidth * gateFarWidth, config.maxHeight * 2 * gateFarHeight).es();
	g.ss(2).s(lineColor).dr(config.maxHeight * borderGap, config.maxHeight * (0.5 - gateNearHeight), config.maxWidth * gateNearWidth, config.maxHeight * 2 * gateNearHeight).es();
	g.ss(2).s(lineColor).dr(config.maxWidth * (1 - gateNearWidth) - config.maxHeight * (borderGap), config.maxHeight * (0.5 - gateNearHeight), config.maxWidth * gateNearWidth, config.maxHeight * 2 * gateNearHeight).es();

	g.ss(2).s(lineColor).f("rgba(255, 255, 255, 0.5)").dr(config.maxHeight * borderGap - config.maxWidth * gateWidth, config.maxHeight * (0.5 - gateHeight), config.maxWidth * gateWidth,	config.maxHeight * 2 * gateHeight).ef().es();
	g.ss(2).s(lineColor).f("rgba(255, 255, 255, 0.5)").dr(config.maxWidth - config.maxHeight * (borderGap), config.maxHeight * (0.5 - gateHeight), config.maxWidth * gateWidth,	config.maxHeight * 2 * gateHeight).ef().es();


	stage.addChild(field);
}

function prettifyDisplay(players) {
	let minDist = 2.1 * config.playerRadius; // TODO ???? (some space please!!!)
	for (let a in players){
		let pa = players[a];
		for (let b in players){		// start from index a and keep going to the end (prevent redundant calculations)
			let pb = players[b];

			if (pa != pb) {
				let ta = pa.getTarget();
				let tb = pb.getTarget();

				let dist = pa.targetDistanceTo(pb);
				if (dist < minDist) {
					let angle = Math.atan2(tb.y - ta.y, tb.x - ta.x);
					tb.x = ta.x + minDist * Math.cos(angle);
					tb.y = ta.y + minDist * Math.sin(angle);
					pb.targetX = tb.x;
					pb.targetY = tb.y;

				}
				if (dist < config.playerLabelRadius){
					pb.clear = false;
				}
			}

		}
	}
}

function cleanupDeadPlayers(players, updatedKeys, lastHolder) {
	for (let key in players) {
		if (!updatedKeys[key]) {
			players[key].kill();
			if (lastHolder == players[key]) {
				lastHolder = null;
			}
			delete players[key];
		}
	}
	return lastHolder;
}









