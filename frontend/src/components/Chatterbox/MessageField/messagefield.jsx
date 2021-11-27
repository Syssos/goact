import React, { Component } from "react";
import "./messagefield.css";

class MessageField extends Component {
  render() {
    return (
      <div className="MessageField">
        <input onKeyDown={this.props.send} placeholder="Send a message..."/>
      </div>
    );
  };
}

export default MessageField;