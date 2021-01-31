import React from 'react';
import {
    AuthenticationDetails,
    CognitoUser,
    CognitoUserPool,
    CognitoUserSession,
    GetSessionOptions,
    ICognitoUserPoolData,
    ISignUpResult,
} from 'amazon-cognito-identity-js';
import { useHistory } from 'react-router-dom';

//#region User Pool
const poolData: ICognitoUserPoolData = {
    UserPoolId: process.env.REACT_APP_USER_POOL_ID || "",
    ClientId: process.env.REACT_APP_CLIENT_ID || "",
}

export const Pool = new CognitoUserPool(poolData)
//#endregion User Pool

//#region Context Types
export type AuthenticateResponse = {
    session: CognitoUserSession;
    userConfirmationNecessary?: boolean;
}

export interface IAccount {
    isAuthenticated: boolean;
    isLoading: boolean;

    authenticate: (email: string, password: string) => Promise<AuthenticateResponse>;
    register: (email: string, password: string) => Promise<ISignUpResult>;
    logout: () => void;
    getSession: () => Promise<CognitoUserSession>;
    requestResetPassword: (email: string) => Promise<unknown>;
    resetPassword: (token: string, password: string) => Promise<unknown>;
}
//#endregion Context Types

//#region Context Definition
export const AccountContext = React.createContext<IAccount>({} as IAccount)
AccountContext.displayName = "AccountContext"
export const useAccount = () => React.useContext(AccountContext)
//#endregion Context Definition

//#region Provider
export type AccountProps = {
    children: React.ReactNode;
}

export const Account = (props: AccountProps) => {
    const [isAuthenticated, setIsAuthenticated] = React.useState<boolean>(false)
    const [isLoading, setIsLoading] = React.useState<boolean>(true)

    const history = useHistory()

    //#region Methods
    const authenticate = (email: string, password: string) =>
        new Promise<AuthenticateResponse>((resolve, reject) => {
            const user = new CognitoUser({ Username: email, Pool })
            const authDetails =
                new AuthenticationDetails({ Username: email, Password: password })

            user.authenticateUser(authDetails, {
                onSuccess: (session, userConfirmationNecessary) => {
                    setIsAuthenticated(true)
                    setIsLoading(false)
                    resolve({ session, userConfirmationNecessary })
                },
                onFailure: err => {
                    setIsAuthenticated(false)
                    setIsLoading(false)
                    reject(err)
                },
                newPasswordRequired: (_userAttributes, requiredAttributes) => {
                    console.warn("New password required")
                    if (requiredAttributes && Array.isArray(requiredAttributes) && requiredAttributes.length !== 0) {
                        console.warn("RequiredAttributes:", requiredAttributes)
                    }
                    setIsAuthenticated(false)
                    setIsLoading(false)
                    reject("new password required")
                }
            })
        })

    const register = (email: string, password: string) =>
        new Promise<ISignUpResult>((resolve, reject) => {
            Pool.signUp(email, password, [], [], (err, result) => {
                if (err) {
                    setIsAuthenticated(false)
                    setIsLoading(false)
                    reject(err)
                } else if (result) {
                    setIsAuthenticated(true)
                    setIsLoading(false)
                    resolve(result)
                } else {
                    setIsAuthenticated(false)
                    setIsLoading(false)
                    reject("failed to register user: empty error")
                }
            })
        })

    const logout = () => {
        const user = Pool.getCurrentUser()
        if (user) {
            localStorage.clear()
            sessionStorage.clear()
            user.signOut()
            setIsAuthenticated(false)
            setIsLoading(false)
            history.push("/")
        }
    }

    const getSession = React.useCallback(() =>
        new Promise<CognitoUserSession>((resolve, reject) => {
            const user = Pool.getCurrentUser()
            if (user) {
                user.getSession((err: Error | null, session: CognitoUserSession) => {
                    if (err) {
                        reject(err)
                    } else {
                        resolve(session)
                    }
                }, {} as GetSessionOptions)
            } else {
                reject("no current user")
            }
        })
        , [])

    const requestResetPassword = (email: string): Promise<unknown> =>
        new Promise((resolve, reject) => {
            if (true) {
                reject("implement")
            }
            resolve("implement")
        })

    const resetPassword = (token: string, password: string): Promise<unknown> =>
        new Promise((resolve, reject) => {
            if (true) {
                reject("implement")
            }
            resolve("implement")
        })
    //#endregion Methods

    //#region Effects
    React.useEffect(() => {
        getSession()
            .then(session => {
                setIsAuthenticated(session.isValid())
                setIsLoading(false)
            })
            .catch(() => {
                setIsAuthenticated(false)
                setIsLoading(false)
            })
    }, [getSession])
    //#endregion Effects

    return (
        <AccountContext.Provider
            value={{
                isAuthenticated,
                isLoading,
                authenticate,
                register,
                logout,
                getSession,
                requestResetPassword,
                resetPassword,
            }}
        >
            { props.children}
        </AccountContext.Provider>
    )
}
//#endregion Provider
