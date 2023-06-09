import React, {Component} from 'react';
import './App.css';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import TextField from 'material-ui/TextField';
import RaisedButton from 'material-ui/RaisedButton';
import Paper from 'material-ui/Paper';
import TagResult from "./components/TagResult";

const style = {
    marginLeft: 12,
    align: 'right',
};

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            sentence: '',
            tags: undefined,
            description: ''
        };
    };

    analyzeSentence() {
        const GO_WEBAPP_URL ='http://localhost:8081/';
        fetch(GO_WEBAPP_URL, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({sentence: this.textField.getValue()})
        })
            .then(response => response.json())
            .then(data => this.setState(data));
    }

    onEnterPress = e => {
        if (e.key === 'Enter') {
            this.analyzeSentence();
        }
    };

    render() {
        const tagResultComponent = this.state.tags !== undefined ?
        <TagResult tags={this.state.tags}/> :
        null;


        return (
            <MuiThemeProvider>
                <div className="centerize">
                    <Paper zDepth={1} className="content">
                        <h2>Tag Builder</h2>
                        <br></br>
                        <TextField ref={ref => this.textField = ref} onKeyUp={this.onEnterPress.bind(this)}
                                   hintText="Type your sentence." multiLine />
                        <RaisedButton  label="Send" style={style} onClick={this.analyzeSentence.bind(this)}/>
                        {tagResultComponent}
                    </Paper>
                </div>
            </MuiThemeProvider>
        );
    }
}

export default App;