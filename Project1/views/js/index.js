
var indexObj;

function Index() {
	this.restful = new Restful();
}

Index.prototype.Initial = function() {
	console.log('index page initial...');
	indexObj.restful.GetTest();
}

window.onload = function () {
	indexObj = new Index();
	indexObj.Initial();
};

