import React from 'react';
import { useHistory } from 'react-router-dom';

import { useAccount } from '../../contexts/Account';
import { RouteBase } from '../Base';

export const Register = () => {
    const [email, setEmail] = React.useState<string>("")
    const [password, setPassword] = React.useState<string>("")
    const [confPassword, setConfPassword] = React.useState<string>("")

    const { register } = useAccount()
    const history = useHistory()

    const onSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()

        if (password !== confPassword || password.trim() === "") return console.error("passwords don't match or password is empty")

        register(email, password)
            .then(result => {
                sessionStorage.setItem(
                    "userConfirmed",
                    result.userConfirmed
                        ? "true"
                        : "false")
                history.push("/")
            })
            .catch(err => console.error(err))
    }

    return (
        <RouteBase>
            <h2>Register:</h2>
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

                <input
                    value={confPassword}
                    placeholder="Confirm Password"
                    onChange={event => setConfPassword(event.target.value)}
                />

                <button type="submit">Register</button>
            </form>
        </RouteBase>
    )
}
