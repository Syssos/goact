import React, { Component } from "react";
import "./BodyLayout.scss";

import Chat from './Chat/Chat';
import Members from './Sidebar/Members/Members';

class BodyLayout extends Component {
  render() {
    return (
      <div className="BodyLayout">
        <Members />
        <Chat
          chatHistory={this.props.chatHistory}
          send={this.props.send} 
        />
      </div>
      );
  };
}

export default BodyLayout;