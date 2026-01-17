

const Button = (props) => {
    const { children, className, type, onClick } = props;
    return (
        <button
            type={type}
            onClick={onClick}
            className={`
                text-white
                px-4 py-1.5 rounded-full
                bg-gradient-to-r from-blue-600 to-purple-600 
                hover:from-blue-500 hover:to-purple-500 
                shadow-lg hover:shadow-blue-500/25
                transition-all duration-200 ease-in-out transform
                hover:scale-[98%]
                ${className}
            `}
        >
            {children}
        </button>
    )
}

export default Button