import React, { Component } from "react";
import { connect, sendMsg } from "./api";

import ChatHistory from './ChatHistory';
import MessageField from './MessageField';
import LoginForm from './LoginForm';

import "./chatterbox.css";


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

class Chatterbox extends Component {
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
    if(!getCookie("token")) return (<div className="Chatterbox"><LoginForm /></div>);

    return (
      <div className="Chatterbox">
        <ChatHistory chatHistory={this.state.chatHistory} />
        <MessageField send={this.send} />
      </div>
      );
  };
}

export default Chatterbox;