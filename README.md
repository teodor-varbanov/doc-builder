# Simple DocBuilder for Ansible Playbooks

The app reads through Ansible playbooks, takes the comments and some other keys (such as block names and when: conditions) and finally puts them through an Ollama hosted LLM that generates documenttion based on the provided information.
We leverage the LLM's strengths in natural language completion and Go's portability, combined with Ollama's ease of use within Go. 

The app is still in early development, though. Since not all results are 100% reliable, I'll need to implement a hybrid solution, utilizing both Templating and LLMs for greater precision (as well as... a lot of other things). 

