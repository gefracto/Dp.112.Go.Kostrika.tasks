angular.module('golang')
	.factory('tasks', function ($http) {
		this.tasks = {
			task1: {
				width: 8,
				height: 4,
				symbol: "#"
			},
			task2: [{
				width: 8,
				height: 5
			}, {
				width: 6,
				height: 9
			}]
		};
	    return {
	    	tasks: this.tasks,
	        runTask : function (id) {
	            return $http.post('/task/' + id, this.tasks[id]);
	        },

	        runAll: function (id) {
	        	return $http.post('/tasks', this.tasks);
	        }
	    }
	});