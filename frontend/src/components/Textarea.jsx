

const TextArea = (props) => {
   const {
      placeholder,
      value,
      onChange,
      label,
      className,
      defaultValue } = props;
   return (
      <div className="mt-4">
         <label htmlFor={label} className="block text-secondary text-lg">{label}</label>
         <textarea
            id={label}
            className={`w-full mt-4 p-2 text-lg border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-primary ${className}`}
            type="text"
            placeholder={placeholder}
            value={value}
            onChange={onChange}
            rows="6"
            defaultValue={defaultValue}
         ></textarea>
      </div>
   )
}

export default TextArea;