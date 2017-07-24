angular.module('golang')
	.constant('DEFAULTS', function() {
  		return {
			1: {
				"width": 10,
				"height": 10,
				"symbol": "*"
			}, 
			2: [{
				"width": 5,
				"height": 5
			}, {
				"width": 10,
				"height": 10
			}]
		};
	}());
