import React from 'react';

import { useAccount } from '../../contexts/Account';

export const Settings = () => {
    const { isAuthenticated, isLoading } = useAccount()

    return !isLoading && isAuthenticated
        ? (
            <div>
                <h2>Settings:</h2>
            </div>
        )
        : <></>
}
