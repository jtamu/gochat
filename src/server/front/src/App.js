import React, {Component} from 'react';
import logo from './logo.svg';
import './App.css';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider'
import TextField from '@material-ui/core/TextField'
import {Card, CardActions, CardHeader, CardText} from 'material-ui/Card'
import {List, ListItem} from 'material-ui/List'
import DropDownMenu from 'material-ui/DropDownMenu'
import MenuItem from 'material-ui/MenuItem'
import RefreshIndicator from 'material-ui/RefreshIndicator'
import Button from '@material-ui/core/Button'
import Icon from '@material-ui/core/Icon'
import Table from '@material-ui/core/Table'
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

class App extends Component {
  constructor(props) {
    super(props)
    this.state = {
      ws: null,
    }
    this.initSocket()
  }

  initSocket() {
    this.state.ws = new WebSocket(`ws://localhost:81/rooms/1/messages/ws`)
  }

  render() {
    return (
      <div className="App" style={{marginLeft: 30, width: 1000}}>
        <MenuArea />
        <MessagesArea ws={this.state.ws} />
        <CommentArea ws={this.state.ws} />
      </div>
    );
  }
}

class MenuArea extends Component {
  constructor(props) {
    super(props)
  }
  render() {
    return (
        // [TODO]
        <Paper zDepth={2}>
          メニューエリア
        </Paper>
    )
  }
}

class MessagesArea extends Component {
  constructor(props) {
    super(props)
    this.state = {
      messages: []
    }
    this.initMessages()
  }

  initMessages() {
    const delay = (mSec) => new Promise(
      (resolve) => setTimeout(resolve, mSec)
    )

    fetch(`http://localhost:81/rooms/1/messages`)
    .then((response) => response.json())
    .then((json) => {
      delay(700)
      .then(() => this.setState({
        messages: json
      }))
    })
    .catch((response) => {
      console.log('** error **', response)
    })

    this.props.ws.onmessage = (msgevent) => {
      console.log("** received **")
      var msg = [JSON.parse(msgevent.data)]
      this.setState({
        messages: this.state.messages.concat(msg)
      })
    }
  }

  render() {
    return (
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ユーザID</TableCell>
              <TableCell>メッセージ</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {this.state.messages.map((message, idx) => <Message key={idx} message={message} />)}
          </TableBody>
        </Table>
      </TableContainer>
    )
  }
}

const Message = (props) => {
  return (
    <TableRow>
      <TableCell>{props.message.user_id}</TableCell>
      <TableCell>{props.message.content}</TableCell>
    </TableRow>
  )
}

class CommentArea extends Component {
  constructor(props) {
    super(props)
    this.state = {
      commentText: ""
    }
  }
  changeText(event) {
    this.setState({commentText: event.target.value})
  }
  comment() {
    var message = {
      content: this.state.commentText
    }
    this.props.ws.send(JSON.stringify(message))
    this.setState({commentText: ""})
    console.log("** send **")
  }
  handleSubmit(event) {
    event.preventDefault()
  }
  keyDown(event) {
    if (event.keyCode == 13) {
      this.comment()
    }
  }
  render() {
    return (
      <Paper zDepth={2}>
        <form onSubmit={this.handleSubmit}>
          <TextField
            label="コメント"
            value={this.state.commentText}
            onChange={(e) => {this.changeText(e)}}
            onKeyDown={(e) => {this.keyDown(e)}}
          />
          <Button
            label="Send"
            variant="contained"
            color="primary"
            endIcon={<Icon>send</Icon>}
            onClick={() => {this.comment()}}
          />
        </form>
      </Paper>
    )
  }
}

export default App;
