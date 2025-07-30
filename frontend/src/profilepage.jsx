import React, { useState, useEffect } from 'react';

const ProfilePage = () => {
  const [username, setUsername] = useState('');
  const [nickname, setNickname] = useState('');
  const [main, setMain] = useState('');
  const [bio, setBio] = useState('');
  const [profilePic, setProfilePic] = useState('');
  const [editing, setEditing] = useState(false);

  const winrate = '75%'; // Static for now

  // Load user data (can be replaced with API fetch)
  useEffect(() => {
    const storedUser = {
      username: localStorage.getItem('username') || '',
      nickname: localStorage.getItem('nickname') || '',
      main: localStorage.getItem('main') || '',
      bio: localStorage.getItem('bio') || '',
      profilePic: localStorage.getItem('profilePic') || '',
    };

    setUsername(storedUser.username);
    setNickname(storedUser.nickname);
    setMain(storedUser.main);
    setBio(storedUser.bio);
    setProfilePic(storedUser.profilePic || 'default-pfp.png');
  }, []);

  // Handle image upload
  const handleImageUpload = (e) => {
    const file = e.target.files?.[0];
    if (!file) return;

    const reader = new FileReader();
    reader.onloadend = () => {
      const dataUrl = reader.result;
      setProfilePic(dataUrl);
      localStorage.setItem('profilePic', dataUrl); // Replace with backend call later
    };
    reader.readAsDataURL(file);
  };

  // Save profile
  const saveChanges = () => {
    localStorage.setItem('username', username);
    localStorage.setItem('nickname', nickname);
    localStorage.setItem('main', main);
    localStorage.setItem('bio', bio);
    setEditing(false);

    // Replace with API call:
    // fetch('/api/updateProfile', { method: 'POST', body: JSON.stringify({ username, nickname, main, bio }) });
  };

  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center py-10">
      <div className="w-full max-w-sm bg-white p-6 rounded-2xl shadow-lg text-center">
        <img
          src={profilePic}
          alt="Profile"
          className="w-24 h-24 mx-auto rounded-full object-cover mb-2"
        />
        {nickname && <h1 className="text-lg italic text-gray-600">"{nickname}"</h1>}
        <h2 className="text-2xl font-semibold">{username || 'Username'}</h2>

        <p className="mt-2"><strong>Main:</strong> {main || 'None'}</p>
        <p><strong>Winrate:</strong> {winrate}</p>

        <div className="mt-3 text-left">
          <p className="font-semibold">Bio:</p>
          <p className="text-sm text-gray-700 whitespace-pre-wrap">
            {bio || 'No bio yet.'}
          </p>
        </div>

        <button
          className="mt-5 bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded-lg"
          onClick={() => setEditing(!editing)}
        >
          {editing ? 'Cancel' : 'Edit Profile'}
        </button>

        {editing && (
          <div className="mt-4 space-y-3 text-left">
            <input
              type="text"
              placeholder="Username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="w-full px-3 py-2 border rounded-lg"
            />
            <input
              type="text"
              placeholder="Nickname"
              value={nickname}
              onChange={(e) => setNickname(e.target.value)}
              className="w-full px-3 py-2 border rounded-lg"
            />
            <input
              type="text"
              placeholder="Main"
              value={main}
              onChange={(e) => setMain(e.target.value)}
              className="w-full px-3 py-2 border rounded-lg"
            />
            <textarea
              placeholder="Bio..."
              value={bio}
              onChange={(e) => setBio(e.target.value)}
              className="w-full px-3 py-2 border rounded-lg resize-none h-24"
            />
            <input
              type="file"
              accept="image/*"
              onChange={handleImageUpload}
              className="w-full"
            />
            <button
              onClick={saveChanges}
              className="w-full bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-lg"
            >
              Save Changes
            </button>
          </div>
        )}
      </div>
    </div>
  );
};

export default ProfilePage;
