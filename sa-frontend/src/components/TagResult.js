import React, {Component} from 'react';
import PropTypes from 'prop-types';

class TagResult extends Component {

    propTypes = {
        tags: PropTypes.array.isRequired,
        description: PropTypes.string.isRequired
    };

    render() {
        const tagsList = this.props.tags.map((tag, index) => <li key={index}>{tag}</li>);
    
        return (
            <div>
                <h3>Tags:</h3>
                <ul>{tagsList}</ul>
                <div>Description: {this.props.description}</div>
            </div>
        );
    }
    
}

export default TagResult;