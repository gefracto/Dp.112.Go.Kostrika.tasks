angular.module('golang')
	.controller('TasksCtrl', function ($scope, tasks) {
		$scope.tasks = tasks.tasks;
		$scope.data = "undefined";
		$scope.error = 'fdsfsdfs';

		$scope.runTask = function (id) {
			console.log('id');
			tasks.runTask(id).then(function (data) {
				console.log(data);
				$scope.data = data.data.response;
				$scope.error = data.data.error;
			});
		};

		$scope.runAll = function () {
			tasks.runAll().then(function (data) {
				console.log(data);
				$scope.data = data.data.response;
				$scope.error = data.data.error;
			});
		}
	});