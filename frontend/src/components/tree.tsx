import {
  ChevronRightIcon,
  Columns,
  Database,
  FileIcon,
  FolderTree,
  Table,
} from "lucide-react";
import {
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
} from "./ui/sidebar";
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "./ui/collapsible";
import { GetChildren } from "../../wailsjs/go/main/App";
import { useEffect, useState } from "react";

export function TreeBuilder({ connId }: { connId: string }) {
  const [tree, setTree] = useState<TreeNode[]>([]);

  useEffect(() => {
    async function loadRoot() {
      const rootId = `conn:${connId}`;
      const children = await GetChildren(rootId);

      setTree(children);
    }

    loadRoot();
  }, [connId]);

  async function handleExpand(node: TreeNode) {
    if (node.children) return;

    updateNode(node.id, { isLoading: true });

    const children = await GetChildren(node.id);

    updateNode(node.id, {
      children,
      isLoading: false,
    });
  }

  function updateNode(id: string, updates: Partial<TreeNode>) {
    function recursive(nodes: TreeNode[]): TreeNode[] {
      return nodes.map((n) => {
        if (n.id === id) {
          return { ...n, ...updates };
        }
        if (n.children) {
          return { ...n, children: recursive(n.children) };
        }
        return n;
      });
    }

    setTree((prev) => recursive(prev));
  }

  return (
    <SidebarMenu>
      {tree.map((node) => (
        <Tree key={node.id} node={node} onExpand={handleExpand} />
      ))}
    </SidebarMenu>
  );
}

type TreeNode = {
  id: string;
  name: string;
  type: string;
  metadata: string;
  hasChildren: boolean;
  children?: TreeNode[];
  isLoading?: boolean;
};

function Tree({
  node,
  onExpand,
}: {
  node: TreeNode;
  onExpand: (node: TreeNode) => void;
}) {
  if (!node.hasChildren) {
    return (
      <SidebarMenuButton>
        <Columns />
        {node.name}
        <p className="text-muted-foreground">{node.metadata}</p>
      </SidebarMenuButton>
    );
  }

  return (
    <SidebarMenuItem>
      <Collapsible
        onOpenChange={() => onExpand(node)}
        className="group/collapsible"
      >
        <CollapsibleTrigger asChild>
          <SidebarMenuButton>
            <ChevronRightIcon className="transition-transform group-data-[state=open]/collapsible:rotate-90" />
            {getIconByType(node.type)}
            {node.name}
          </SidebarMenuButton>
        </CollapsibleTrigger>

        <CollapsibleContent>
          <SidebarMenuSub className="m-0.5">
            {node.isLoading && <div>Loading...</div>}

            {node.children?.map((child) => (
              <Tree key={child.id} node={child} onExpand={onExpand} />
            ))}
          </SidebarMenuSub>
        </CollapsibleContent>
      </Collapsible>
    </SidebarMenuItem>
  );
}

function getIconByType(type: string) {
  switch (type) {
    case "database":
      return <Database size={16} />;
    case "schema":
      return <FolderTree size={16} />;
    case "table":
      return <Table size={16} />;
    case "column":
      return <Columns size={16} />;
    default:
      return <FileIcon size={16} />;
  }
}
