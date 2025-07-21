import React from 'react'
import './PublicChat.css';

function PublicChat() {
    return (

        <div className="main-section">

            <div className="public-chat">
                
                <div className="vert-wrap">

                    {/* Title Text */}
                    <div className = "title-text">Public Chat</div>

                    {/* Chat Log */}
                    <div className="chat-log">

                        <div className="chat-header"></div>

                        {/* Message 1 */}

                        <div className="message-header">

                            <img src="/enzoPFP.png" alt="enzo pfp" />

                            <div className="name-date-spacing"> 

                                <div className="name-and-date">enzo</div>
                                <div className="name-and-date">6:14 PM</div>

                            </div>

                        </div>

                        <div className="text-message">Hey guys</div>

                        <div className="divider"></div>

                        {/* Message 2 */}

                        <div className="message-header">

                            <img src="/calebPFP.png" alt="enzo pfp" />

                            <div className="name-date-spacing"> 

                                <div className="name-and-date">caleb</div>
                                <div className="name-and-date">6:17 PM</div>

                            </div>

                        </div>

                        <div className="text-message">Looking for someone to practice w/</div>

                        <div className="divider"></div>

                        

                    </div>

                </div>

            </div>

        </div>
  )
}

export default PublicChat
