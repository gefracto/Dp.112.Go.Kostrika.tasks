angular.module('golang')
	.constant('DEFAULTS', function() {
  		return {
			1: {
				"width": 8,
				"height": 4,
				"symbol": "#"
			}, 
			2: [{
				"width": 8,
				"height": 5
			}, {
				"width": 6,
				"height": 9
			}]
		};
	}());