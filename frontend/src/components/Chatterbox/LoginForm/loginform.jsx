import React, { useState } from 'react';
import axios from 'axios';
import "./loginform.css";


const LoginForm = () => {
	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');
	const [error, setError] = useState('');

	const handleSubmit = async (e) => {
		e.preventDefault();

		const authObject = { 'Username': username, 'Password': password }
		console.log(authObject)

		try {
			await axios.post('http://localhost:8080/validate', authObject, {withCredentials: true});
			localStorage.setItem('User', username);

			window.location.reload();
		} catch (error) {
	 			// Wrong sign-in information
				setError('Oops incorrect credentials...')
		}
	}

	return (
		<div className="wrapper">
			<div className="form">
				<h1 className="title">Chatterbox</h1>
				<form onSubmit={handleSubmit}>
					<input type="text" value={username} onChange={(e) => setUsername(e.target.value)} className="input" placeholder="Username" required />
					<input type="password" value={password} onChange={(e) => setPassword(e.target.value)} className="input" placeholder="Password" required />
					<div align="center">
						<button type="submit" className="button">
							<span>Start Chatting</span>
						</button>
					</div>
					<h2 className="error">{error}</h2>
				</form>
			</div>
		</div>
	);
}

export default LoginForm;