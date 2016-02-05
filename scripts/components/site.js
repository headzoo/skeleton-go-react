'use strict';

var React  = require('react');
var Reflux = require('reflux');
var Navbar = require('./navbar');

var Component = React.createClass({
    
    render: function () {
        return (
            <div className="container">
                <Navbar />
                {this.props.children}
            </div>
        )
    }
});

module.exports = Component;
