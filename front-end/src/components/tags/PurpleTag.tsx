import { Typography, Tag } from "antd";

const { Text } = Typography;

export const PurpleTag = ({ label, closable = false, onClose = () => {} } : 
{ label : string, closable? : boolean,  onClose?: () => void }
) => {
  const onPreventMouseDown = (event: React.MouseEvent<HTMLSpanElement>) => {
    event.preventDefault();
    event.stopPropagation();
  };
  return (
    <Tag
        color="purple"
        onMouseDown={onPreventMouseDown}
        closable={closable}
        onClose={onClose}
        style={{ marginRight: 3, marginTop: 1, marginBottom: 1 }}
    >
      <Text>{label}</Text>
    </Tag>
  );
};