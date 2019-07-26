function DashboardService($log, $http) {
    var users = [{
        Names: ['dopware'],
        Id: 'beyond',
    }];

    $log = $log.getInstance("DashboardService");
    $log.debug("instanceOf() ");

    // Promise-based API
    return {
        loadNodes: function() {
            $log.debug("loadNodes()");
            return $http.get('/api/nodes');
        },
        loadContainers: function(node) {
            $log.debug("loadContainers()");
            return $http.get('/api/nodes/containers?node=' + node);
        }
    };
}


export default ['$log', '$http', DashboardService];
