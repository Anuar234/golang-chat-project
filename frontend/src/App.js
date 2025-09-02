import React, {Component} from 'react';
import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';
import ChatInput from './components/ChatInput/ChatInput';
import {connect, sendMsg} from './api';
import './App.css';

class App extends Component{
  constructor(props){
    super(props);
    this.state = {
      chatHistory: []  // Исправлено: строчная буква
    }
  }

  componentDidMount(){
    connect((msg) => {
      console.log("New Message:", msg)
      this.setState(prevState => ({
        chatHistory: [...prevState.chatHistory, msg]
      }))
      console.log(this.state);
    })
  } 

  // Добавлен метод send для отправки сообщений
  send = (event) => {
    if(event.key === 'Enter') {
      sendMsg(event.target.value);
      event.target.value = ""; // Очищаем поле ввода
    }
  }

  render() {
    return(
      <div className='App'>
        <Header/>
        <ChatHistory chatHistory={this.state.chatHistory}/>
        <ChatInput send={this.send}/>
      </div>
    );
  }
}

export default App;