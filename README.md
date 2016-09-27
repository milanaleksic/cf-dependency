# CloudFormation Dependency Graph

This is a trivial tool that, given list of cloudformation template files as program arguments, provides a simple
DOT-driven PNG dependency graph (DOT compiler from GraphViz which should be installed).

It only tracks explicit dependencies via `DependsOn` between resources. This is done on purpose because of 
behavior of CF that (sometimes?) ignores implicit dependencies.

run as:

    go install github.com/milanaleksic/cf-dependency
    cf-dependency archives.template archives-api.template | dot -Tpng -o graph.png
