import React, { Component } from "react";
import "./Chat.scss";

import ChatHistory from './ChatHistory/ChatHistory';
import ChatInput from './ChatInput/ChatInput';

class Chat extends Component {
  render() {
    return (
      <div className="Chat">
        <ChatHistory chatHistory={this.props.chatHistory} />
        <ChatInput send={this.props.send} />
      </div>
      );
  };
}

export default Chat;