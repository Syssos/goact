import React, { Component } from "react";
import "./BodyLayout.scss";

import Chat from './Chat/Chat';

class BodyLayout extends Component {
  render() {
    return (
      <div className="BodyLayout">
        <Chat
          chatHistory={this.props.chatHistory}
          send={this.props.send} 
        />
      </div>
      );
  };
}

export default BodyLayout;