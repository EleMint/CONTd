import React from 'react';
import styled from 'styled-components';

import { useAccount } from '../contexts/Account';
import { DraggableVerticalSeparator } from './DraggableVerticalSeparator';

type HamburgerProps = {
    onClick: (event: React.MouseEvent<HTMLDivElement, MouseEvent>) => void;
    open: boolean;
    width: number;
}

const Hamburger = (props: HamburgerProps) => (
    <HamburgerBase
        onClick={props.onClick}
        open={props.open}
        width={props.width}
    >
        <HamburgerBar className={`bar1 ${props.open ? "change" : ""}`} />
        <HamburgerBar className={`bar2 ${props.open ? "change" : ""}`} />
        <HamburgerBar className={`bar3 ${props.open ? "change" : ""}`} />
    </HamburgerBase>
)

const HamburgerBase = styled.div<HamburgerProps>`
    position: fixed;
    height: 35px;
    width: 35px;
    left: ${props => props.width}%;
    margin-left: 10px;
    top: 100px;
    z-index: 1000;
    display: inline-block;
    cursor: pointer;

    .change.bar1 {
        -webkit-transform: rotate(-45deg) translate(-9px, 6px);
        transform: rotate(-45deg) translate(-9px, 6px);
    }

    .change.bar2 {
        opacity: 0;
    }

    .change.bar3 {
        -webkit-transform: rotate(45deg) translate(-8px, -8px);
        transform: rotate(45deg) translate(-8px, -8px);
    }
`

const HamburgerBar = styled.div`
    width: 35px;
    height: 5px;
    background-color: #333;
    margin: 6px 0;
    transition: 0.5s;
`

export const Menu = () => {
    const [open, setOpen] = React.useState<boolean>(false)
    const [width, setWidth] = React.useState<number>(0)

    const { isAuthenticated, isLoading } = useAccount()

    const onHamburgerClick = () => {
        setOpen(!open)
        setWidth(!open ? 30 : 0)
    }

    return !isLoading && isAuthenticated
        ? (
            <>
                <Hamburger
                    onClick={onHamburgerClick}
                    open={open}
                    width={width}
                />
                { (open && <DraggableVerticalSeparator
                    width={width}
                    setWidth={setWidth}
                    onDoubleClick={() => open && onHamburgerClick()}
                />) || <></>}
                <Base width={width}>

                </Base>
            </>
        )
        : <></>
}

const Base = styled.div<{ width: number }>`
    position: absolute;
    top: 0;
    left: 0;
    z-index: 1000;
    top: 100px;
    height: calc(100% - 100px);
    width: ${props => props.width}%;
    background: red;
    overflow: scroll;
`
