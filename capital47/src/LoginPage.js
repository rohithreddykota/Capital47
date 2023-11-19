// LoginPage.js
import React, { useState } from 'react';
import './LoginPage.css'; // Make sure to adjust the path based on your file structure
import { useNavigate } from 'react-router-dom';



function LoginPage() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [isLoginMode, setLoginMode] = useState(true);

    console.log('I am login')

    const history = useNavigate(); // Access the history object

    const handleToggleMode = () => {
        setLoginMode((prevMode) => !prevMode);
    };

    const handleNavigation = () => {
        // Assuming you have routes set up for '/currency-sending' in your App component
        history.push('/currency-sending');
    };

    const handleSubmit = (e) => {
        e.preventDefault();

        // Add your login or account creation logic here
        if (isLoginMode) {
            console.log('Login submitted:', { username, password });
            // Add logic for login
        } else {
            console.log('Account created:', { username, password });
            // Add logic for account creation
        }

        // Navigate to the currency sending page
        handleNavigation();
    };

    return (
        <div className="login-container">
            <h2>{isLoginMode ? 'Login' : 'Create Account'}</h2>
            <form onSubmit={handleSubmit}>
                <label>
                    Username:
                    <input type="text" value={username} onChange={(e) => setUsername(e.target.value)} />
                </label>
                <br />
                <label>
                    Password:
                    <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
                </label>
                <br />
                <button type="submit">{isLoginMode ? 'Login' : 'Create Account'}</button>
            </form>
            <p onClick={handleToggleMode} className="toggle-mode">
                {isLoginMode ? 'Create an account' : 'Already have an account? Login'}
            </p>
        </div>
    );
}

export default LoginPage;
