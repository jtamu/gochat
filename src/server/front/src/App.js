import React, {Component} from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props)
    this.getMessages()
  }

  getMessages() {
    fetch(`http://localhost:81/rooms/1/messages`)
    .then((response) => response.json())
    .then((json) => {
      console.log(json)
    })
    .catch((response) => {
      this.setState({loading: false})
      console.log('** error **', response)
    })
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.js</code> and save to reload.
          </p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
        </header>
      </div>
    );
  }
}

export default App;
