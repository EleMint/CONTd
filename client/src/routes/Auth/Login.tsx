import React from 'react';
import { useHistory } from 'react-router-dom';

import { useAccount } from '../../contexts/Account';
import { RouteBase } from '../Base';

export const Login = () => {
    const [email, setEmail] = React.useState<string>("")
    const [password, setPassword] = React.useState<string>("")

    const { authenticate } = useAccount()
    const history = useHistory()

    const onSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()

        if (email.trim() === "" || password.trim() === "") return

        authenticate(email, password)
            .then(resp => {
                sessionStorage.setItem("userConfirmed", resp.userConfirmationNecessary ? "true" : "false")
                history.push("/")
            })
            .catch(err => console.error("Failed to login:", err))
    }

    return (
        <RouteBase>
            <h2>Login:</h2>
            <form onSubmit={onSubmit}>
                <input
                    value={email}
                    placeholder="Email"
                    onChange={event => setEmail(event.target.value)}
                />

                <input
                    value={password}
                    placeholder="Password"
                    onChange={event => setPassword(event.target.value)}
                />

                <button type="submit">Login</button>
            </form>
        </RouteBase>
    )
}
