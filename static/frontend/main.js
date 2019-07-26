import { ExternalLogger } from 'utils/LogDecorator';
import dashboard from 'dashboard/Dashboard';

let $log = new ExternalLogger();
    $log = $log.getInstance( "BOOTSTRAP" );
    $log.debug( "Configuring 'main' module" );

export default angular.module('main', [dashboard]).name;
