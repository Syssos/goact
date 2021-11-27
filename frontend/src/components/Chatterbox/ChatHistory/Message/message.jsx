import React, { Component } from "react";
import "./message.css";

class Message extends Component {
  constructor(props) {
    super(props);
    let temp = JSON.parse(this.props.message);
    this.state = {
      message: temp
    };
  }

  render() {
    if (localStorage.getItem('User') === this.state.message.user) {
      return <div className="Message me">{this.state.message.user} - {this.state.message.body}</div>
    }
    return <div className="Message other">{this.state.message.user} - {this.state.message.body}</div>;
  }
}

export default Message;