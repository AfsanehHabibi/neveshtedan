import React, { useState } from 'react';
import { useMutation, useQuery, gql } from '@apollo/client';
import { useParams } from 'react-router-dom';

const GET_WRITING_TEMPLATE = gql`
  query GetWritingTemplate($id: Int!) {
    template: writingTemplate(id: $id) {
      title
      id
      fields
    }
  }
`;

const CREATE_WRITING_ENTRY_MUTATION = gql`
  mutation CreateWritingEntryMutation($input: NewWritingEntry!) {
    createWritingEntry(input: $input) {
      id
      userId
      templateId
      fields {
        name
        value
      }
    }
  }
`;

const NewWriting = () => {
  const { id } = useParams();
  const { loading, error, data } = useQuery(GET_WRITING_TEMPLATE, {
    variables: { id: parseInt(id, 10) },
  });

  const [formState, setFormState] = useState({
    // Initialize formState with default values or values from the template
    value1: '',
    value2: '',
    // Add more fields as needed
  });

  const [createWritingEntry] = useMutation(CREATE_WRITING_ENTRY_MUTATION);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  const template = data.template;

  const handleSubmit = (e) => {
    e.preventDefault();

    // Use template.fields to dynamically create fields in the input object
    let newEntry = {
      fields: template.fields.map((fieldName) => ({
        name: fieldName,
        value: formState[fieldName.toLowerCase()], // Assuming lowercase field names
      })),
      templateId: template.id,
    };

    // Retrieve the token from local storage
    const token = localStorage.getItem('token');

    console.log(token)
    createWritingEntry({
      variables: { input: newEntry },
      context: {
        headers: {
          Authorization: token ? `${token}` : '',
        },
      },
    })
      .then((result) => console.log(result))
      .catch((err) => console.log(err));
  };

  return (
    <div>
      <h2>{template.title}</h2>
      <form onSubmit={handleSubmit}>
        <div className="flex flex-column mt3">
          {template.fields.map((fieldName) => (
            <input
              key={fieldName}
              className="mb2"
              value={formState[fieldName.toLowerCase()]}
              onChange={(e) =>
                setFormState({
                  ...formState,
                  [fieldName.toLowerCase()]: e.target.value,
                })
              }
              type="text"
              placeholder={fieldName}
            />
          ))}
        </div>
        <button type="submit">Submit</button>
      </form>
    </div>
  );
};

export default NewWriting;
