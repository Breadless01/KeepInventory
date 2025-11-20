export function Button({ children, className = "", ...props }) {
  return (
    <button
      className={"ki-btn " + className}
      type="button"
      {...props}
    >
      {children}
    </button>
  );
}
