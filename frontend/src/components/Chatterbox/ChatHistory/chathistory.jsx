import React, { Component } from "react";
import "./chathistory.css";

import Message from './Message';

class ChatHistory extends Component {
  render() {
    const messages = this.props.chatHistory.map(msg => <Message message={msg.data} />);
    
    return (
      <div className="ChatHistory">
        {messages}
      </div>
      );
  };
}

export default ChatHistory;