import React from 'react';
import { RouteComponentProps } from 'react-router-dom';
import styled from 'styled-components';

import { useAccount } from '../../contexts/Account';
import { RouteBase } from '../Base';

export const RequestResetPassword = (props: RouteComponentProps) => {
    const [email, setEmail] = React.useState<string>("")

    const { requestResetPassword } = useAccount()

    const onSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()

        if (email.trim() === "") return console.error("email cannot be empty")

        requestResetPassword(email)
            .then(resp => console.log(resp))
            .catch(err => console.error(err))
    }

    return (
        <RouteBase>
            <h2>Request Reset Password:</h2>
            <form onSubmit={onSubmit}>
                <input
                    value={email}
                    placeholder="Email"
                    onChange={event => setEmail(event.target.value)}
                />

                <button type="submit">Request</button>
            </form>
        </RouteBase>
    )
}
