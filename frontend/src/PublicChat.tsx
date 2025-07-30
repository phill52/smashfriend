import React from 'react';
import './PublicChat.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPaperPlane } from '@fortawesome/free-regular-svg-icons';
import { useEffect, useRef, useState } from 'react';

interface Message {
    id: number;
    name: string;
    time: string;
    text: string;
    profileImg: string;
}

const ROOM_ID = "0";
const USERNAME = "andre";

function PublicChat() {
    const [messages, setMessages] = useState<Message[]>([]);
    const [input, setInput] = useState<string>("");
    const bottomRef = useRef<HTMLDivElement | null>(null);

    useEffect(() => {
        bottomRef.current?.scrollIntoView({ behavior: "smooth" });
    }, [messages]);

    const handleSend = () => {
        if (!input.trim()) return;

    const mockMessage: Message = {
        id: Date.now(),
        name: USERNAME,
        time: new Date().toLocaleTimeString([], { hour: 'numeric', minute: '2-digit' }),
        text: input,
        profileImg: "/enzoPFP.png",
    };

        setMessages(prev => [...prev, mockMessage]);
        setInput("");

    };

    return (
        <div className="main-section">

            <div className="public-chat">

                <div className="vert-wrap">

                    <div className="title-text">Public Chat</div>

                    <div className="chat-log">

                        <div className="chat-header" />

                        {messages.map((msg) => (

                            <div key={msg.id}>

                                <div className="message-header">

                                    <img src={msg.profileImg} alt={`${msg.name} pfp`} />

                                    <div className="name-date-spacing">

                                        <div className="name-and-date">{msg.name}</div>

                                        <div className="name-and-date">{msg.time}</div>

                                    </div>
                                </div>

                                <div className="text-message">{msg.text}</div>

                                <div className="divider" />

                            </div>

                        ))}

                        <div ref={bottomRef} />

                    </div>

                    <div className="message-component">

                        <input

                            type="text"
                            placeholder="Message Here"
                            value={input}
                            onChange={(e) => setInput(e.target.value)}
                            className="message-input"
                            onKeyDown={(e) => e.key === "Enter" && handleSend()}

                        />

                        <FontAwesomeIcon icon={faPaperPlane} className="messageIcon" onClick={handleSend} />

                    </div>
                </div>
            </div>
        </div>
    );
}

export default PublicChat;
