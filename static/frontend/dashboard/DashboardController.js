/**
 * Main App Controller for the Angular Material Starter App
 * @param $scope
 * @param $mdSidenav
 * @param avatarsService
 * @constructor
 */
function DashboardController( dashboardService, $mdSidenav, $mdBottomSheet, $log ) {

  $log = $log.getInstance( "SessionController" );
  $log.debug( "instanceOf() ");

  var self = this;

  self.selected     = null;
  self.containers   = [];
  self.nodes        = [ ];
  self.selectNode   = selectNode;
  self.toggleList   = toggleNodesList;
//  self.terminal     = terminal;

  // Load all nodes

  dashboardService
        .loadNodes()
        .then(function(payload) {
          self.nodes    = [].concat(payload.data);

          self.selected = payload.data[0];

          dashboardService
            .loadContainers(self.selected.status.addresses[0].address)
            .then(function(payload) {
                self.containers = payload.data
            });
        });

  // *********************************
  // Internal methods
  // *********************************

  /**
   * Hide or Show the 'left' sideNav area
   */
  function toggleNodesList() {
    $log.debug( "toggleNodesList() ");
    $mdSidenav('left').toggle();
  }

  /**
   * Show Containers
   * @param menuId
   */
  function selectNode ( node ) {
    $log.debug( "selectNode( {metadata.name} ) ", node);

    self.selected = angular.isNumber(node) ? $scope.nodes[node] : node;

    dashboardService
        .loadContainers(node.status.addresses[0].address)
        .then(function(payload) {
            self.containers = payload.data
        });

    self.toggleList();
  }

//  function terminal(container) {
//    $log.debug("terminal")
//
//    var node = self.selected;
//    $log.debug("node: {metadata.name} ip: {address} ", node.status.addresses[0]);
//    $log.debug("containerId {Id}", container);
//  }

}

export default [
    'dashboardService', '$mdSidenav', '$mdBottomSheet', '$log',
    DashboardController
  ];

