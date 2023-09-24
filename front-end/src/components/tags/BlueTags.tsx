import { Typography, Tag } from "antd";
import type { CustomTagProps } from 'rc-select/lib/BaseSelect';

const { Text } = Typography;

export const BlueTags = (props: CustomTagProps) => {
  const { label, closable, onClose } = props;
  const onPreventMouseDown = (event: React.MouseEvent<HTMLSpanElement>) => {
    event.preventDefault();
    event.stopPropagation();
  };
  return (
    <Tag
      color="blue"
      onMouseDown={onPreventMouseDown}
      closable={closable}
      onClose={onClose}
      style={{ marginRight: 3, marginTop: 1, marginBottom: 1 }}
    >
      <Text style={{color: '#0958d9'}}>{label}</Text>
    </Tag>
  );
};