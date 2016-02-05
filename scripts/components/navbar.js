'use strict';

var React  = require('react');
var Reflux = require('reflux');
var activeComponent = require('react-router-active-component');

var Component = React.createClass({

    render: function () {
        var NavLink = activeComponent('li');
    
        return (
            <nav className="navbar navbar-default">
                <div className="container-fluid">
                    <div className="navbar-header">
                        <a className="navbar-brand" href="#">Website</a>
                    </div>
                    <div id="navbar">
                        <ul className="nav navbar-nav">
                            <NavLink to="/" onlyActiveOnIndex>Home</NavLink>
                            <NavLink to="/about">About</NavLink>
                            <NavLink to="/contact">Contact</NavLink>
                        </ul>
                    </div>
                </div>
            </nav>
        )
    }
});

module.exports = Component;
