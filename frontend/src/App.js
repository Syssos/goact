// App.js
import React, { Component } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";

// Import our new component from it's relative path
import ChatHistory from './components/ChatHistory/ChatHistory';
import ChatInput from './components/ChatInput/ChatInput';
import Header from './components/Header/Header';
import LoginForm from './components/LoginForm'

function getCookie(name) {
    var dc = document.cookie;
    var prefix = name + "=";
    var begin = dc.indexOf("; " + prefix);
    if (begin === -1) {
        begin = dc.indexOf(prefix);
        if (begin !== 0) return null;
    }
    else
    {
        begin += 2;
        var end = document.cookie.indexOf(";", begin);
        if (end === -1) {
        end = dc.length;
        }
    }
    // because unescape has been deprecated, replaced with decodeURI
    //return unescape(dc.substring(begin + prefix.length, end));
    return decodeURI(dc.substring(begin + prefix.length, end));
}

class App extends Component {
  constructor(props) {
    super(props);
    // Creating single chat history instance
    this.state = {
      chatHistory: []
    }
  }

  // Connecting to websocket
  componentDidMount() {
    connect((msg) => {
      // Adding msg to chat history
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory, msg]
      }))
    });
  }

  send(event) {
    // Checking if enter was pressed
    if(event.keyCode === 13) {
      sendMsg(event.target.value);
      event.target.value = "";
    }
  }

  render() {
    if(!getCookie("token")) return <LoginForm />
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={this.send} />
      </div>
    );
  }
}

export default App;