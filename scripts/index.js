var React                = require('react');
var ReactDOM             = require('react-dom');
var ReactRouter          = require('react-router');
var createBrowserHistory = require('history/lib/createBrowserHistory');
var Router               = ReactRouter.Router;
var Route                = ReactRouter.Route;
var Link                 = ReactRouter.Link;
var IndexRoute           = ReactRouter.IndexRoute;

var Site    = require('./components/site');
var Home    = require('./components/home');
var About   = require('./components/about');
var Contact = require('./components/contact');
var NoMatch = require('./components/nomatch');

$(function () {
    ReactDOM.render((
            <Router history={createBrowserHistory()}>
                <Route path="/" component={Site}>
                    <IndexRoute component={Home}/>
                    <Route path="about" component={About}/>
                    <Route path="contact" component={Contact}/>
                    <Route path="*" component={NoMatch}/>
                </Route>
            </Router>),
        document.getElementById("mount")
    );
});