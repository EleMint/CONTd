import React from 'react';
import styled from 'styled-components';

const VerticalSeparator = styled.div<{ width: number }>`
    position: fixed;
    top: 100px;
    left: ${props => props.width}%;
    height: calc(100% - 100px);
    width: 4px;
    z-index: 1001;
    background-color: grey;
    cursor: col-resize;
`

export const DraggableVerticalSeparator = (props: {
    width: number,
    setWidth: React.Dispatch<React.SetStateAction<number>>,
    onDoubleClick: () => void
}) => {
    const closeDragEvent = function (this: GlobalEventHandlers, event: MouseEvent) {
        document.onmouseup = null
        document.onmousemove = null
    }

    const elementDrag = function (this: GlobalEventHandlers, event: MouseEvent) {
        event.preventDefault()

        if (event.pageX !== 0) {
            props.setWidth((event.pageX - 2) * 100 / window.outerWidth)
        }
    }

    const dragMouseDown = (event: React.MouseEvent<HTMLDivElement, MouseEvent>) => {
        event.preventDefault()

        if (event.pageX !== 0) {
            props.setWidth((event.pageX - 2) * 100 / window.outerWidth)
        }
        document.onmouseup = closeDragEvent
        document.onmousemove = elementDrag
    }

    return (
        <VerticalSeparator
            width={props.width}
            onDoubleClick={props.onDoubleClick}
            onMouseDown={dragMouseDown}
        />
    )
}