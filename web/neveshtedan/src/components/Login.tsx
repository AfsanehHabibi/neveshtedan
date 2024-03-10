import React, { useState } from 'react';
import { useMutation, gql } from '@apollo/client';

const LOGIN_MUTATION = gql`
  mutation Login($input: Login!) {
    login(input: $input)
  }
`;

const Login = ({ onLogin }) => {
  const [formState, setFormState] = useState({
    username: '',
    password: '',
  });

  const [login, { loading, error }] = useMutation(LOGIN_MUTATION);

  const handleSubmit = async (e) => {
    e.preventDefault();

    const userInput = {
      username: formState.username,
      password: formState.password,
    };

    try {
      const result = await login({
        variables: { input: userInput },
      });

      const token = result.data.login;

      // Handle the token, e.g., save to local storage
      localStorage.setItem('token', token);

      // Notify the parent component about successful login
      onLogin();
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div>
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <div className="flex flex-column mt3">
          <input
            className="mb2"
            value={formState.username}
            onChange={(e) => setFormState({ ...formState, username: e.target.value })}
            type="text"
            placeholder="Username"
          />
          <input
            className="mb2"
            value={formState.password}
            onChange={(e) => setFormState({ ...formState, password: e.target.value })}
            type="password"
            placeholder="Password"
          />
        </div>
        <button type="submit" disabled={loading}>
          {loading ? 'Logging in...' : 'Login'}
        </button>
        {error && <p>Error: {error.message}</p>}
      </form>
    </div>
  );
};

export default Login;
