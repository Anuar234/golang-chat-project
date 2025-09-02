import React, {Component} from 'react';
import './ChatHistory.scss';
import Message from '../Message/Message';

class ChatHistory extends Component{
    render(){
        console.log("ChatHistory props:", this.props.chatHistory);
        
        // Исправлено: сохраняем результат map в переменную
        const messages = this.props.chatHistory.map((msg, index) => (
            <Message key={index} message={JSON.stringify(msg)}/>
        ));

        return(
            <div className='ChatHistory'>
                <h2>Chat History</h2>
                {messages}
            </div>
        )
    }
}

// Добавлен экспорт
export default ChatHistory;