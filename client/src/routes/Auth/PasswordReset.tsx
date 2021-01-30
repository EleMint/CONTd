import React from 'react';
import { RouteComponentProps } from 'react-router-dom';

import { useAccount } from '../../contexts/Account';
import { RouteBase } from '../Base';

const getToken = (): string => {
    return ""
}

export const ResetPassword = (props: RouteComponentProps) => {
    const token = getToken()
    const [password, setPassword] = React.useState<string>("")
    const [confPassword, setConfPassword] = React.useState<string>("")

    const { resetPassword } = useAccount()

    const onSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()

        if (password !== confPassword || password.trim() === "") return console.error("passwords don't match or password is empty")

        resetPassword(token, password)
            .then(result => console.log(result))
            .catch(err => console.error(err))
    }

    return (
        <RouteBase>
            <h2>Reset Password:</h2>
            <form onSubmit={onSubmit}>
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

                <button type="submit">Reset</button>
            </form>
        </RouteBase>
    )
}
