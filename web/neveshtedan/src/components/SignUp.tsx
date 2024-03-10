import React, { useState } from 'react';
import { useMutation, gql } from '@apollo/client';

const SIGNUP_MUTATION = gql`
  mutation CreateUser($input: NewUser!) {
    createUser(input: $input)
  }
`;

const SignUp = () => {
  const [formState, setFormState] = useState({
    username: '',
    password: '',
  });

  const [createUser] = useMutation(SIGNUP_MUTATION);

  const handleSubmit = (e) => {
    e.preventDefault();

    const newUser = {
      username: formState.username,
      password: formState.password,
    };

    createUser({
      variables: { input: newUser },
    })
      .then((result) => {
        const token = result.data.createUser;
        // Save the token to local storage
        localStorage.setItem('token', token);

        // Handle successful sign-up, e.g., redirect to another page
      })
      .catch((err) => console.log(err));
  };

  return (
    <div>
      <h2>Sign Up</h2>
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
        <button type="submit">Sign Up</button>
      </form>
    </div>
  );
};

export default SignUp;
