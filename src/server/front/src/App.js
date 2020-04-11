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
        <form>
          <input type="text" />
        </form>
      </div>
    );
  }
}

export default App;
