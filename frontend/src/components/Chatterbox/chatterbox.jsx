import React, { Component } from "react";
import { connect, sendMsg } from "./api";

import ChatHistory from './ChatHistory';
import MessageField from './MessageField';

import "./chatterbox.css";

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
    return (
      <div className="Chatterbox">
        <ChatHistory chatHistory={this.state.chatHistory} />
        <MessageField send={this.send} />
      </div>
      );
  };
}

export default Chatterbox;