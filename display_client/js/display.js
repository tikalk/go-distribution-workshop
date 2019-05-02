// TODO remove dead players from display

let Player = class {
	constructor(model, config) {
		this.config = config;

		this.container = new createjs.Container();

		this.hl = new createjs.Shape();
		this.hl.graphics.beginFill("yellow").drawCircle(0, 0, config.playerRadius * 2);
		this.hl.alpha = 0.5;
		this.container.addChild(this.hl);

		this.graphic = new createjs.Shape();
		this.graphic.graphics.beginFill(model.team_id).drawCircle(0, 0, config.playerRadius);
		this.container.addChild(this.graphic);

		this.hl.visible = false;

		let text = new createjs.Text(model.item_label, "15px Arial", "#333");
		text.x = config.playerRadius * 1.1;
		text.y = 7.5;
		text.textBaseline = "alphabetic";
		this.container.addChild(text);

		this.update(model, true);

		config.stage.addChild(this.container);
	}

	update(model, hard){
		if(model) {
			this.targetX = this.config.maxWidth * model.x / 100;
			this.targetY = this.config.maxHeight * model.y / 100;
		}

		if (hard) {
			this.container.x = this.targetX;
			this.container.y = this.targetY;
		}
		else {
			this.container.x += (this.targetX - this.container.x) * 0.2;
			this.container.y += (this.targetY - this.container.y) * 0.2;
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
}

let Ball = class {
	constructor(model, config) {
		this.config = config;

		this.graphic = new createjs.Shape();
		this.graphic.graphics.beginFill("black").drawCircle(0, 0, config.ballRadius);

		if(!model){
			model = {x: 0, y : 0}
		}
		this.update(model, true);
		config.stage.addChild(this.graphic);
	}

	update(model, hard){
		if(model) {
			this.targetX = this.config.maxWidth * model.x / 100;
			this.targetY = this.config.maxHeight * model.y / 100;
		}

		if (true) {
			this.graphic.x = this.targetX;
			this.graphic.y = this.targetY;
		}
		else {
			this.graphic.x += (this.targetX - this.graphic.x) * 0.2;
			this.graphic.y += (this.targetY - this.graphic.y) * 0.2;
		}
	}
}

async function init() {
	let canvas = document.getElementById("main_canvas");
	canvas.width = window.innerWidth;
	canvas.height = window.innerHeight;

	let config = {
		maxWidth: window.innerWidth,
		maxHeight: window.innerHeight,
		playerRadius: 10,
		displayInterval: 100,
		animationInterval: 25,
		ballRadius: 5,
	}

	let stage = new createjs.Stage("main_canvas");

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

		for (let key in players){
			if (!updatedKeys[key]){
				players[key].kill();
				if (lastHolder == players[key]) {
					lastHolder = null;
				}
				delete players[key];
			}
		}

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





