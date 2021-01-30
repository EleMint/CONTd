import React from 'react';

//#region Context Types
export interface IAPI {
    get: (url: string, headers?: Record<string, string>) => Promise<unknown>;
}
//#endregion Context Types

//#region Context Definition
export const APIContext = React.createContext<IAPI>({} as IAPI)
APIContext.displayName = "APIContext"
export const useAPI = () => React.useContext(APIContext)
//#endregion Context Definition

//#region Provider
export type APIProps = {
    children: React.ReactNode;
}

export const API = (props: APIProps) => {
    //#region Methods
    const get = (url: string, headers?: Record<string, string>) =>
        new Promise((resolve, reject) => {
            if (true) {
                resolve("implement")
            }
            reject("implement")
        })
    //#endregion Methods

    return (
        <APIContext.Provider
            value={{
                get
            }}
        >
            { props.children}
        </APIContext.Provider>
    )
}
//#endregion Provider
