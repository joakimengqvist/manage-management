import { Typography, Tag } from "antd";
import type { CustomTagProps } from 'rc-select/lib/BaseSelect';

const { Text } = Typography;

export const PurpleTags = (props: CustomTagProps) => {
  const { label, closable, onClose } = props;
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